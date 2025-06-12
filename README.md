Dataset 1: DOTA v1
Installation procedure- 
1. Go to https://github.com/ultralytics/assets/releases/download/v0.0.0/DOTAv1.zip
2. Your download will automatically start (File Size - 1.99GB)
3. Place the files in this heirarchy:\n
    parent\n
    ├── ultralytics\,
    └── datasets\n
        └── dota1  ← downloads here (2GB)\n
4. To train the dataset efficiently we have to split each image into resolution of 1024x1024. Run DOTAv1-Prep.py to prepare the data for        training automatically




