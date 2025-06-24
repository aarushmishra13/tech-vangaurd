*All commands assume a Bash-like shell (macOS/Linux, WSL, or PowerShell with `bash` syntax).*

_____________________________________________________________________________________________

Drone/Overhead/ISR Datasets:
 
**Dataset 1: UAVDT**

```bash
# 1. Install/Upgrade Dataset Tools
pip install --upgrade dataset-tools

# 2. Download UAVDT
python UAVDT.py
```
*Run UAVDT.py in your **working directory** to download the dataset in the proper heirarchy.*

**Dataset 2: VisDrone**

```bash
# 1. Install/Upgrade Ultralytics
pip install --upgrade ultralytics

# 2. Download VisDrone
yolo train data=VisDrone.yaml
```
*The dataset will automatically download to /datasets/VisDrone and the annotations will be converted to YOLO format.*

**Dataset 3: DOTAv1**

```bash
# 1. Install/Upgrade Ultralytics
pip install --upgrade ultralytics

# 2. Download DOTAv1
yolo train data=DOTAv1.yaml
```
*The dataset will automatically download to /datasets/DOTA-split and the annotations will be converted to YOLO format.*

_____________________________________________________________________________________________

RF Signal/Waveform Datasets:

**Dataset 1: Noisy Drone RF (ZHAW)**

```bash
# 1. Download
pip install kaggle torchaudio matplotlib pillow torch
mkdir -p datasets/NoisyDroneRF
kaggle datasets download sgluege/noisy-drone-rf-signal-classification -p datasets/NoisyDroneRF
unzip datasets/NoisyDroneRF/*.zip -d datasets/NoisyDroneRF

# 2. Convert IQ → spectrogram PNGs (224 × 224) & split
python convert_noisydrone.py --root datasets/NoisyDroneRF --img_size 224

# 3. Copy YAML to project root (adjust path inside YAML if you move the folder)
cp datasets/NoisyDroneRF/NoisyDroneRF.yaml .
```

**Dataset 2: TorchSig (synthetic)**

```bash
# 1. Install + generate a 10 k-sample narrowband set
pip install torchsig torchaudio matplotlib pillow
python generate_torchsig_dataset.py --root datasets/TorchSig --samples 10000

# 2. Copy YAML
cp datasets/TorchSig/TorchSig.yaml .
```

**Dataset 3: RadioML 2016.10A**

```bash
# 1. Download
pip install kaggle numpy torchaudio matplotlib pillow
mkdir -p datasets/RadioML2016
kaggle datasets download nolasthitnotomorrow/radioml2016-deepsigcom -p datasets/RadioML2016
unzip datasets/RadioML2016/*.zip -d datasets/RadioML2016

# 2. Convert
python convert_radioml2016.py --root datasets/RadioML2016 --img_size 224

# 3. Copy YAML
cp datasets/RadioML2016/RadioML2016.yaml .
```
_____________________________________________________________________________________________
Dataset 1: DOTAv2

# 1. Install Git-LFS (large-file support)
pip install git-lfs
git lfs install                          # one-time setup

# 2. Download (≈2 GB)
git clone https://huggingface.co/datasets/satellite-image-deep-learning/DOTAv2

# 3. Convert & tile to YOLO-OBB (1024 × 1024 crops)
python converters/dota2yolo_obb.py \
       --dota_root datasets/DOTAv2 \
       --out_root datasets/DOTAv2/labels



Dataset 2: BigEarthNet (Sentinel-1 + Sentinel-2)

# 1. Manual download (~59 GB optical + 51 GB SAR)
#    https://bigearth.net/#downloads  (grab both S2 and S1 zips)

# 2. Unzip to the expected folders
mkdir -p datasets/BigEarthNet-S2  && unzip BigEarthNet-S2.zip  -d datasets/BigEarthNet-S2
mkdir -p datasets/BigEarthNet-S1  && unzip BigEarthNet-S1.zip  -d datasets/BigEarthNet-S1



Dataset 3: Military Aircraft Detection

# 1. Install Kaggle CLI
pip install kaggle pandas

# 2. Download & unzip  (≈11 GB)
mkdir -p datasets/MilitaryAircraftDetection
kaggle datasets download a2015003713/militaryaircraftdetectiondataset -p datasets/MilitaryAircraftDetection
unzip datasets/MilitaryAircraftDetection/*.zip -d datasets/MilitaryAircraftDetection

# 3. Convert CSV → YOLO
python converters/mad_csv2yolo.py \
       --csv     datasets/MilitaryAircraftDetection/labels_with_split.csv \
       --img_dir datasets/MilitaryAircraftDetection/dataset \
       --out_dir datasets/MilitaryAircraftDetection/labels



Dataset 4: MAR20 Military Aircraft Recognition

# 1. Download (≈1.2 GB)
mkdir -p datasets/MAR20
kaggle datasets download khlaifiabilel/military-aircraft-recognition-dataset -p datasets/MAR20
unzip datasets/MAR20/*.zip -d datasets/MAR20

# 2. Convert Pascal-VOC XML → YOLO
python converters/voc2yolo.py \
       --ann_dir datasets/MAR20/Annotations/Horizontal\\ Bounding\\ Box \
       --img_dir datasets/MAR20/JPEGImages \
       --out_dir datasets/MAR20/labels

  



Dataset 5: HRSID (High-Res SAR Ship)

# 1. Clone repo (≈3 GB)
git clone https://github.com/chaozhong2010/HRSID datasets/HRSID

# 2. Convert COCO JSON → YOLO
python converters/coco2yolo.py \
       --json      datasets/HRSID/annotations/instances_train.json \
       --img_root  datasets/HRSID/images/train \
       --out_dir   datasets/HRSID/labels/train
_____________________________________________________________________________________________

Automotive Radar Datasets:

**Dataset 1: Heriot-Watt RADIATE Dataset**
```bash
# Step 1
git clone https://github.com/marcelsheeny/radiate_sdk.git

# Step 2
cd radiate_sdk

# Step 3
pip install -r requirements.txt
```

**Dataset 2: The RadarScenes data set**
```bash
# Step 1 Install kagglehub
pip install kagglehub

# Step 2 Run the python script to load data in pandas dataframe
import kagglehub
from kagglehub import KaggleDatasetAdapter

# Set the path to the file you'd like to load
file_path = ""

# Load the latest version
df = kagglehub.load_dataset(
  KaggleDatasetAdapter.PANDAS,
  "aleksandrdubrovin/the-radarscenes-data-set",
  file_path,
  # Provide any additional arguments like 
  # sql_query or pandas_kwargs. See the 
  # documenation for more information:
  # https://github.com/Kaggle/kagglehub/blob/main/README.md#kaggledatasetadapterpandas
)
```


**Dataset 3: Austin Radar Traffic Dataset**
```bash
# Step 1 Install kagglehub
pip install kagglehub

# Step 2 Run the python script to load data in pandas dataframe
import kagglehub
from kagglehub import KaggleDatasetAdapter

# Set the path to the file you'd like to load
file_path = ""

# Load the latest version
df = kagglehub.load_dataset(
  KaggleDatasetAdapter.PANDAS,
  "vinayshanbhag/radar-traffic-data",
  file_path,
  # Provide any additional arguments like 
  # sql_query or pandas_kwargs. See the 
  # documenation for more information:
  # https://github.com/Kaggle/kagglehub/blob/main/README.md#kaggledatasetadapterpandas
)

```

