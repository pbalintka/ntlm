import os

for root, dirs, files in os.walk("."):
    path = root.split(os.sep)
    if len(path)>2 and os.path.isdir(root) and path[2].startswith("."):
        npath = ['.', path[1]+path[2]]
        nroot = "/".join(npath)
        print("{} -> {}".format(root, nroot))
        os.rename(root, nroot)
