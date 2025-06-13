Dataset 1: UAVDT
Installation procedure-

1. Run pip install --upgrade dataset-tools in the terminal.
2. Run UAVDT.py in your working directory to download the dataset in the proper heirarchy.

Dataset 2: VisDrone
Installation procedure-

1. Go to terminal and run: yolo train data=VisDrone.yaml
2. The dataset will automatically download to /datasets/VisDrone and the annotations will be converted to YOLO format.

Dataset 3: DOTAv1
Installation procedure-

1. Place DOTAv1.yaml in your working directory
2. Go to terminal and run: yolo train data=DOTAv1.yaml
3. The dataset will automatically download to /datasets/DOTA-split and the annotations will be converted to YOLO format.

Dataset 1-Aerial/SAR: DOTAv2 
Installation procedure-

1. pip install git-lfs
2. git lfs install
3. git clone https://huggingface.co/datasets/satellite-image-deep-learning/DOTAv2   # 2 GB, contains DOTAv2.zip

