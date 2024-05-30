import argparse
import os
import subprocess


item_list = os.listdir(".")
ignoreFolders = ["build", "common"]
folders = []
for item in item_list:
    if not item.startswith('.') and os.path.isdir(os.path.join(".", item)) and item not in ignoreFolders:
        folders.append(item)


def tidy():
    folders.insert(0, "common")
    for folder in folders:
        os.chdir(folder)
        result = subprocess.run(
            "go mod tidy", shell=True)
        if result.returncode != 0:
            raise Exception(f"tidy {folder} error")
        os.chdir('..')


def build():
    os.makedirs("build", exist_ok=True)
    for folder in folders:
        os.chdir(folder)
        result = subprocess.run(
            "go build -o ../build", shell=True)
        if result.returncode != 0:
            raise Exception(f"build {folder} error")
        os.chdir('..')


parser = argparse.ArgumentParser(description='Help to build')
parser.add_argument('action', choices=[
                    'tidy', 'build'], help='Choose one of the actions')
args = parser.parse_args()
if args.action == 'tidy':
    tidy()
elif args.action == 'build':
    build()
else:
    raise Exception("unknow arg")
