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

--------

# RF Signal / Waveform Datasets
*All commands assume a Bash-like shell (macOS/Linux, WSL, or PowerShell with `bash` syntax).*

## Dataset 1 — Noisy Drone RF (ZHAW)

```bash
# 1 Download
pip install kaggle torchaudio matplotlib pillow torch
mkdir -p datasets/NoisyDroneRF
kaggle datasets download sgluege/noisy-drone-rf-signal-classification -p datasets/NoisyDroneRF
unzip datasets/NoisyDroneRF/*.zip -d datasets/NoisyDroneRF

# 2 Convert IQ → spectrogram PNGs (224 × 224) & split
python convert_noisydrone.py --root datasets/NoisyDroneRF --img_size 224

# 3 Copy YAML to project root (adjust path inside YAML if you move the folder)
cp datasets/NoisyDroneRF/NoisyDroneRF.yaml .
```

## Dataset 2 — TorchSig (synthetic)

```bash
# 1 Install + generate a 10 k-sample narrowband set
pip install torchsig torchaudio matplotlib pillow
python generate_torchsig_dataset.py --root datasets/TorchSig --samples 10000

# 2 Copy YAML
cp datasets/TorchSig/TorchSig.yaml .
```

## Dataset 3 — RadioML 2016.10A

```bash
# 1 Download
pip install kaggle numpy torchaudio matplotlib pillow
mkdir -p datasets/RadioML2016
kaggle datasets download nolasthitnotomorrow/radioml2016-deepsigcom -p datasets/RadioML2016
unzip datasets/RadioML2016/*.zip -d datasets/RadioML2016

# 2 Convert
python convert_radioml2016.py --root datasets/RadioML2016 --img_size 224

# 3 Copy YAML
cp datasets/RadioML2016/RadioML2016.yaml .
```


