func convertTemperature(celsius float64) []float64 {
    kel := celsius + 273.15
    fah := celsius* 1.8 + 32

    return []float64{kel,fah}
}