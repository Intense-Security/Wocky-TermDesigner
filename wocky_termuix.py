##
##          Wocky TermUIX Library/Example Script
##
## @creator: vZy
##
## In order to use this file. include it to your project folder and include '' from file.path import * '' 
## in the file your willing to use these functions in!
##
## Be sure to replace all ' {apikey} ' text to your API KEY!
import os, sys, time, requests

class WockyTermUIX:
    ## method -> img2txt(url)
    ## description -> converting image to asni text and returns the ASNI text URL
    def img2txt(url):
        return requests.get(f"https://wocky.pw/termuix/?key={apikey}&action=img2txt&url={url}").text

    ## method -> Save_ASNI_To_File(filepath, url)
    ## argument note -> enter full filepath to file!
    ## description -> converting image to asni text then saving it to a file
    def GetASNICode(file_path, url):
        asni_text = requests.get(f"https://wocky.pw/termuix/?key={apikey}&action=img2txt&url={url}").text
        try:
            asni_file = open(file_path, "w")
            asni_file.write(asni_text)
            asni_file.close()
            return f"ASNI Text saved to {file_path}"
        except:
            return "[x] Error, Unable to open file!"
    
    ## method -> gradient_fade(color1, color2, text)
    ## description -> converting text into gradient color fade!
    def gradient_fade(color1, color2, text):
        return requests.get(f"https://wocky.pw/termuix/?key={apikey}&action=gradient-fade-save&text={text}&color1={color1}&color2={color2}").text