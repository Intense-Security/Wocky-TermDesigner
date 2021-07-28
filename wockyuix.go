package wocky_wrapper

import (
    "errors"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

func Get_Image2AnsiF(api_key string, writer io.Writer, image_url string) (int, error) {

    var New = http.Client{
        Timeout: 8 * time.Second,
    }

    Resp, error := New.Get("https://wocky.pw/img2txt.php?url=" + url.QueryEscape(image_url) + "&key=" + url.QueryEscape(api_key))
    if error != nil {
        return 0, error
    }

    if Resp.StatusCode != 200 {
        return 0, errors.New("wocky api: returned connection error code: " + strconv.Itoa(Resp.StatusCode))
    }

    Body, error := ioutil.ReadAll(Resp.Body)
    if error != nil {
        return 0, error
    }

    if strings.Contains(string(Body), "[x]") {
        return 0, errors.New(strings.Replace(string(Body), "[x]", "", -1))
    }

    Out := strings.ReplaceAll(string(Body), "URL to your ASNI Code: ", "")

    Resp2nd, error := New.Get(Out)
    if error != nil {
        return 0, error
    }

    if Resp2nd.StatusCode != 200 {
        return 0, errors.New("wocky api: returned connection error code: " + strconv.Itoa(Resp.StatusCode))
    }

    Body2nd, error := ioutil.ReadAll(Resp2nd.Body)
    if error != nil {
        return 0, error
    }

    return writer.Write([]byte(string(Body2nd)))
}

func Get_Image2AnsiString(api_key string, image_url string) (string, error) {
    var New = http.Client{Timeout: 8 * time.Second,}

    Resp, error := New.Get("https://wocky.pw/img2txt.php?url=" + url.QueryEscape(image_url) + "&key=" + url.QueryEscape(api_key))
    if error != nil { return "", error }

    if Resp.StatusCode != 200 { return "", errors.New("wocky api: returned connection error code: " + strconv.Itoa(Resp.StatusCode)) }

    Body, error := ioutil.ReadAll(Resp.Body)
    if error != nil { return "", error }

    if strings.Contains(string(Body), "[x]") { return "", errors.New(strings.Replace(string(Body), "[x]", "", -1)) }

    Out := strings.ReplaceAll(string(Body), "URL to your ASNI Code: ", "")

    Resp2nd, error := New.Get(Out)
    if error != nil { return "", error }

    if Resp2nd.StatusCode != 200 { return "", errors.New("wocky api: returned connection error code: " + strconv.Itoa(Resp.StatusCode)) }

    Body2nd, error := ioutil.ReadAll(Resp2nd.Body)
    if error != nil { return "", error }

    return string(Body2nd), nil
}
