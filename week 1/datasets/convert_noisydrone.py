import torch, argparse, os, matplotlib.pyplot as plt
from torchaudio.transforms import Spectrogram
spec = Spectrogram(n_fft=256, power=2)

def main(root):
    ds = torch.load(f"{root}/dataset.pt")
    classes = ds['classes']            # list comes inside pt file
    for i, (iq, y, snr) in enumerate(zip(ds['data'], ds['labels'], ds['snrs'])):
        split = 'val' if i % 10 == 0 else 'train'
        out_dir = f"{root}/{split}/{classes[y]}"
        os.makedirs(out_dir, exist_ok=True)
        plt.imsave(f"{out_dir}/{i}_{snr}.png", spec(iq).log2().numpy(), format='png')

if __name__ == "__main__":
    ap = argparse.ArgumentParser()
    ap.add_argument("--root", default="datasets/NoisyDroneRF")
    args = ap.parse_args(); main(args.root)
