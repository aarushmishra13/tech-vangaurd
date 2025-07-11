
from pathlib import Path
from PIL import Image
import numpy as np

dataset_dir = Path("classification_dataset/train")  # or val if you want
sizes = []

for class_dir in dataset_dir.iterdir():
    for img_path in class_dir.glob("*.jpg"):
        try:
            with Image.open(img_path) as img:
                sizes.append(img.size)  # (width, height)
        except:
            print(f"⚠️ Skipping broken image: {img_path}")

sizes = np.array(sizes)
avg_width = int(np.median(sizes[:, 0]))
avg_height = int(np.median(sizes[:, 1]))
max_width = int(np.max(sizes[:, 0]))
max_height = int(np.max(sizes[:, 1]))

print(f"📏 Average size: {avg_width}x{avg_height}")
print(f"📐 Max size: {max_width}x{max_height}")
