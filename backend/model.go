package main


type Image struct {
	URL string `json:"url"`
	GCP string `json:"gcp"`
	IMG string `json:"img"`
}

// func toBase64(b []byte) string {
// 	return base64.StdEncoding.EncodeToString(b)
// }

// func fromBase64(s string) ([]byte, error) {
// 	return base64.StdEncoding.DecodeString(s)
// }

// func loadImg(fileName string) (string, error) {
// 	bytes, err := ioutil.ReadFile(fileName)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//imageType := http.DetectContentType(bytes)
// 	//var imgString string

// 	//czy to ma leciec do GCP???
// 	// switch imageType {
// 	// case "image/jpeg":
// 	// 	imgString += "data:image/jpeg;base64," //czy trzeba te znaczniki dodawać przy wyslaniu do GCP??
// 	// case "image/png":
// 	// 	imgString += "data:image/png;base64,"
// 	// }

// 	imgString := toBase64(bytes)

// 	return imgString, err
// }

// func saveImg(imgString string, outFile string) {

// 	img, _ := fromBase64(imgString)
// 	err := ioutil.WriteFile(outFile, img, 0666)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func downloadFile(URL, fileName string) error {
// 	//Get the response bytes from the url
// 	response, err := http.Get(URL)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode != 200 {
// 		return errors.New("received non 200 response code")
// 	}
// 	//Create a empty file
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	//Write the bytes to the fiel
// 	_, err = io.Copy(file, response.Body)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//func fromMain(){
	//stworzyc obiekt database -> connection -> sprawdzic erorr osobna funkcja
	//handler powinine zwracac odpowiednie kody bledu, np. 500 internal serv. err

	/*inFile := "image.png"
	err := downloadFile("https://i.pinimg.com/originals/54/9b/11/549b114dad455ae154295ecad1d05f71.png", inFile)
	if err != nil {
		fmt.Println(err)
	}
	img, err := loadImg(inFile)
	if err != nil {
		fmt.Println(err)
	}

	j, err := json.Marshal(Image{URL: img, GCP: "gcp"}) //key: url | value: {url:"...", GCP:"..."}
	if err != nil {
		fmt.Println(err)
	}

	insertToDB("id", string(j), connection)

	imgFromDB, err := getFromDB("id", connection)
	if err != nil {
		fmt.Println(err)
	}

	var x Image
	err = json.Unmarshal([]byte(imgFromDB.(string)), &x)
	if err != nil {
		fmt.Println(err)
	}

	saveImg(x.URL, "outFile.png")*/
//}