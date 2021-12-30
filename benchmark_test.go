package main

import (
	"fmt"
	"testing"
	"strings"
	"math"
	"io/ioutil"
	"github.com/UniversityTeam/SoftwareEngineeringLab4/engine"
)

func AddNewStringToChange(k int) {
    for i := 0; i < k; i++ {
        input,_  := ioutil.ReadFile("./testFile.txt")
        parts := strings.Fields(string(input))
        replacedString := fmt.Sprintf("%s%s", parts[1], parts[1])
        newCmd := fmt.Sprintf("%s %s %s", parts[0], replacedString, parts[2])
        ioutil.WriteFile("./testFile.txt", []byte(newCmd), 0644)
    }
}

func InitNewFile() {
    ioutil.WriteFile("./testFile.txt", []byte("delete abcdefg a"), 0644)
}

func BenchmarkParse(b *testing.B) {
	for i := 1; i <= 20; i++ {
	    InitNewFile()
	    AddNewStringToChange(i)
	    cmd, _ := ioutil.ReadFile("./testFile.txt")
		b.Run(fmt.Sprintf("%d-length", int(math.Pow(2.0, float64(i)))), func(b *testing.B) {
		    for i := 0; i < b.N; i++ {
		        engine.Parse(string(cmd))
		    }
		})
	}
}