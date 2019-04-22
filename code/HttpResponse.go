func HttpResponse(url string) (string, error) {
	defer tools.MRecover()
	response, err := http.Get(url)
	if err != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"url": url,
			"err": string(err.Error()),
		})
		logger.Error("open url has error")
		return "nil", errors.New("open url has error : " + err.Error())
	}
	defer response.Body.Close()
	rep, err := ioutil.ReadAll(response.Body)
	if err != nil {
		var logger = logrus.WithFields(logrus.Fields{
			"url": url,
			"err": string(err.Error()),
		})
		logger.Error("read body error")
		return "nil", errors.New(err.Error())
	}
	return string(rep), nil
}
