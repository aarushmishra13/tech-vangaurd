import os, argparse, torch, matplotlib.pyplot as plt
from torchsig.datasets.datamodules import NarrowbandDataModule     # :contentReference[oaicite:4]{index=4}
from torchaudio.transforms import Spectrogram
spec = Spectrogram(n_fft=256, power=2)

ap = argparse.ArgumentParser()
ap.add_argument("--root", default="datasets/TorchSig")
ap.add_argument("--samples", type=int, default=10000)
args = ap.parse_args()

dm = NarrowbandDataModule(root=args.root, num_train=args.samples, impaired=True)
dm.prepare_data(); dm.setup()

for split, ds in [("train", dm.train_dataset), ("val", dm.val_dataset)]:
    for i, (iq, y) in enumerate(ds):
        cls = dm.classes[y]
        out_dir = f"{args.root}/{split}/{cls}"; os.makedirs(out_dir, exist_ok=True)
        plt.imsave(f"{out_dir}/{i}.png", spec(iq).log2().numpy(), format='png')
