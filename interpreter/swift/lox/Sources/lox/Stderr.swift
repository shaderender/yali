import Foundation

struct Stderr: TextOutputStream {
    mutating func write(_ string: String) {
        fputs(string, stderr)
    }
}

var standardError = Stderr()
