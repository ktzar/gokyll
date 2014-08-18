#Gokyll

##What's Gokyll?

Golang static page generator.

This project is still in an early development stage. It's not even usable yet.

##Getting started
Run `gokyll` indicating the directory where your project is stored. Gokyll will process it and, if everything is correct, will output the static site in the `_site` subdirectory.

##Configuration
The `config.json` file contains the configuration for your website. There are some keys that must be present in it, whereas others are optional.

##Directories
Directories in your project starting with `_` will not be present in the final generated site. These have different functionalities that will be better described in the next sections. The rest of the directories will be copied to the output folder straight away, without any processing at all. These will usually contain your static style and script files.

##HTML files
All `.html` files present in the main directory will be processed according to the following rules. The processed file will be copied to the output directory.

###Variables
Your HTML files can use the following variables:
- `{.Title}` The title of the site
- `{.Year}` The current year


##Templates
The `_templates` directory contains html files that declare sections that can be overriden by your html files. Your html files need to extend from one template.

##SASS files
TODO - Not yet implemented
Files in the `_sass` directory will be processed




