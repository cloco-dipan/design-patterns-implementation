### Factory Pattern implementation in Go ###


This project demonstrates a simple implementation of the Factory Pattern in Go. The Factory Pattern is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

## Project Structure

- `logger/base.go`: Contains the `Logger` interface and the `LoggerFactory` function which implements the Factory Pattern.
- `logger/console.go`: Contains the `ConsoleLogger` struct which implements the `Logger` interface.
- `logger/file.go`: Contains the `FileLogger` struct which implements the `Logger` interface.
- `app.log`: Example log file used for testing.

The Factory Pattern in this project is used to create instances of different types of loggers (console and file) based on a provided type. Here's how it works:  

1. #### Logger Interface ####
    The Logger interface defines a single method Log that all loggers must implement.  
```type Logger interface {
Log(interface{})
}
```

2. #### ConsoleLogger ####
    The ConsoleLogger struct implements the Logger interface and logs messages to the console.  
```
type ConsoleLogger struct {}
func (cl *ConsoleLogger) Log(msg interface{}) {
// Implementation for logging to console
}
```

3. #### FileLogger ####
    The FileLogger struct (presumably defined in logger/file.go) also implements the Logger interface and logs messages to a file.  

```
type FileLogger struct {}
func (fl *FileLogger) Log(msg interface{}) {
// Implementation for logging to file
}
```

4. #### LoggerFactory #### 
    The LoggerFactory function takes a string parameter loggerType and returns an instance of a logger based on the provided type.  

````
func LoggerFactory(loggerType string) Logger {
    switch loggerType {
        case "console":
            return &ConsoleLogger{}
        case "file":
            return &FileLogger{}
    }
    return nil
}
````

5. #### Main Function ####
    The main function demonstrates the use of the LoggerFactory function to create instances of different loggers and log messages using the common Logger interface.

````
var CL = LoggerFactory("console")
var FL = LoggerFactory("file")

func main() {
    CL.Log("test string log")
    FL.Log("test file log")
}
````
In summary, the Factory Pattern allows the creation of different types of loggers (console and file) through a common interface (Logger). The LoggerFactory function encapsulates the logic of instantiating the appropriate logger based on the provided type, promoting flexibility and scalability in the logging mechanism.

