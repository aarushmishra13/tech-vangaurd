\# Week-4



\## What Did We Do?



This week, every member of our team did their individual work on using the dataset prepared in week-3 to fine-tune a YOLOv8 model.



\## How Did We Do It?



We trained our detection model on the 9 US aircrafts we chose. We used Yolo library's tuner for hyperparameter tuning. Then we used these parameters to train the model. We achieved 95% accuracy on average.



One problem we faced was that the model was labeling the whole image instead of labeling only the aircraft. We are suspecting that this is happening because we used cropped images and used a script to label 100% of the images. 



These errors will be resolved next week.



\## Files and Deliverables



The detection model that is trained on the top 9 US aircrafts is in Faisal's folder. It contains the best hyperparameters, best model, the code and a short pdf for documentation.

