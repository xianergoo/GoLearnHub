import os
import shutil

# 指定文件夹路径
folder_path = "./"

# 遍历文件夹中的文件
for filename in os.listdir(folder_path):
    file_path = os.path.join(folder_path, filename)
    if os.path.isfile(file_path):
        # 获取文件名的前缀
        prefix = filename.split(".", 1)[0]
        if prefix.startswith("p"):
            # 创建文件夹 
            folder_name = os.path.join(folder_path, prefix[:2])
            if not os.path.exists(folder_name):
                os.makedirs(folder_name)
                print(folder_name)
            # 移动文件到对应的文件夹
            shutil.move(file_path, folder_name)