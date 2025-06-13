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

Aerial/SAR Datasets:

**Dataset 1: DOTAv2**

```bash
# Step 1
pip install git-lfs

# Step 2
git lfs install

# Step 3
git clone https://huggingface.co/datasets/satellite-image-deep-learning/DOTAv2   
```  

Aerial/SAR Datasets:

**Dataset 1: DOTAv2**

```bash
# Step 1
pip install git-lfs

# Step 2
git lfs install

# Step 3
git clone https://huggingface.co/datasets/satellite-image-deep-learning/DOTAv2   
```  

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






