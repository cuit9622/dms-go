import argparse
import glob
import multiprocessing
import os
import subprocess


ignoreFolders = ["build", "common"]
item_list = os.listdir(".")
folders: list[str] = []
for item in item_list:
    if not item.startswith('.')\
            and os.path.isdir(os.path.join(".", item))\
            and item not in ignoreFolders:
        folders.append(item)


def executeProcess(folder: str, command: str):
    os.chdir(folder)
    result = subprocess.run(
        command, shell=True)
    if result.returncode != 0:
        raise Exception(f"excute {command} error in {folder}")


def execute(folders: list[str], command: str):
    processes = []
    for folder in folders:
        process = multiprocessing.Process(target=executeProcess, args=(
            folder, command))
        process.start()
        processes.append(process)

    for process in processes:
        process.join()


def replace_text_in_file(file_path, old_text, new_text):
    with open(file_path, 'r+', encoding='UTF-8') as file:
        file_content = file.read()
        modified_content = file_content.replace(old_text, new_text)
        file.seek(0)
        file.write(modified_content)
        file.truncate()


def tidy():
    command = "go mod tidy"
    executeProcess("common", command)
    execute(folders, command)


def build():
    os.makedirs("build", exist_ok=True)
    execute(folders, "go build -o ../build")


def proto():
    os.chdir("common/pb")
    executeProcess(
        ".", "protoc --proto_path=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto")
    fileList = glob.glob("*.pb.go")
    for file in fileList:
        replace_text_in_file(file, ",omitempty", "") #删除omitempty


if __name__ == '__main__':
    multiprocessing.freeze_support()  # 在Windows下必须调用此函数
    parser = argparse.ArgumentParser(description='Help to build')
    parser.add_argument('action', choices=[
                        'tidy', 'build', 'proto'], help='Choose one of the actions')
    args = parser.parse_args()
    if args.action == 'tidy':
        tidy()
    elif args.action == 'build':
        build()
    elif args.action == 'proto':
        proto()
    else:
        raise Exception("unknow arg")
