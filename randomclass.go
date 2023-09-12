package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "time"
)

const timeBuffer = 2 * time.Second // 2-second buffer to avoid self-trigger

func generateNewClass(lastClass string) (string, error) {
    classes := []string{
        "class recon", "class nurse", "class engineer", "class marine",
        "class cyborg", "class arsonist", "class gunner", "class sniper", "class spy",
    }

    // Remove lastClass from the array
    var newClasses []string
    for _, class := range classes {
        if class != lastClass {
            newClasses = append(newClasses, class)
        }
    }

    // Randomly select from the remaining classes
    randomIndex := rand.Intn(len(newClasses))
    chosenClass := newClasses[randomIndex]

    // Write to file
    err := ioutil.WriteFile("randomclass.cfg", []byte(chosenClass), 0644)
    if err != nil {
        return "", err
    }

    return chosenClass, nil
}

func main() {
    fmt.Println("https://github.com/injate/q3wfa-randomclass")
    fmt.Println("Welcome to the Quake 3 WFA Class Randomizer!")
    fmt.Println("To use this with your Quake 3 configuration, you can add 'exec randomclass.cfg' to your autoexec.cfg file.")
    fmt.Println("Alternatively, you can use 'bind <key> exec randomclass.cfg' to execute the class change on demand.")
    fmt.Println("Now generating your classes...")

    rand.Seed(time.Now().Unix())

    lastClass, err := generateNewClass("")
    lastAccessTime := time.Now().Add(-timeBuffer) // Initialize to a past time
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    for {
        time.Sleep(250 * time.Millisecond)

        currentAccessTime, err := getLastAccessTime("randomclass.cfg")
        if err != nil {
            fmt.Println("Error getting file info:", err)
            return
        }

        if currentAccessTime.After(lastAccessTime.Add(timeBuffer)) {
            //fmt.Println("lastAccessTime: ", lastAccessTime)
            lastAccessTime = currentAccessTime

            lastClass, err = generateNewClass(lastClass)
            if err != nil {
                fmt.Println("Error writing to file:", err)
                return
            }
            fmt.Println("Something read randomclass @", lastAccessTime, "- new class generated", lastClass)
        }
    }
}
