/*
               Wocky TermUIX Library/Example Script

@creator: vZy

In order to use this file. include it to your project folder and include '' var ASNI = require("./wocky_termuix.js"); '' 
in the file your willing to use these functions in!

Be sure to replace all ' {apikey} ' text to your API KEY!
*/
var fs = require("fs");
var fetch = require("node-fetch");

function GetASNICode(url) {
    return fetch("https://wocky.pw/img2txt/?key={apikey}&action=img2txt&url=" + url).text
}

function Save_ASNI_To_File(file_path, url) {
    let asni_code = fetch("https://wocky.pw/img2txt/?key={apikey}&action=img2txt&url=" + url).text
    fs.writeFileSync(file_path, asni_code);
    return "ASNI Sequence Text Added Successfully"
}