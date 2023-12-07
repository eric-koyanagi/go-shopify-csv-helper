# Go Shopify CSV Helper
Shopify is a great platform, but it's sometimes suprising that their CSV importer is limited to only 15 MB. 

For users with a lot of products, this forces them to use an API or awkwardly split up a CSV...I'm sure there's a wealth of tools out there to split a Shopify CSV and avoid that 15 MB limit, but I decided to write a utility for this, myself. 

# How to Use
This is a go app, so you'll need to install go. This was done as a coding project, so it expects that you can look into the code yourself. 

- Update the main() endpoint to change the CSV file location. 
- In the `fixCSvSize` module, adjust the `MAX_ROWS_PER_FILE` const to create larger or smaller files. 
- By default, resized files will be prefxied like "resizedCsv_1.csv", "resizedCsv_2.csv", etc. 
- Simply drag and drop the mess of files into Shopify. Fun. 

Note that this only outputs new files if the input file is > 15 MB. Also, it checks the input file to ensure that each column matches the Shopify definition. 

# Creating dummy files 
See https://github.com/eric-koyanagi/shopify-csv-maker