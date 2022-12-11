import Foundation

@main
struct Lox {
    static var hadError = false
    
    static func main() {
        do {
            try runFromArgs()
        } catch {
            print(error, to: &standardError)
            exit(64)
        }
    }
    
    private static func runFromArgs() throws {
        switch Args.get() {
        case let .unrecognized(args):
            print("error: unrecognized arguments: \(args)\n", to: &standardError)
            print(Args.usage(), to: &standardError)
            exit(64)
        case .help:
            print(Args.usage())
        case .prompt:
            try runPrompt()
        case let .run(path):
            try runFile(path: path)
        case .format(_):
            print("error: format is not yet implemented", to: &standardError)
            exit(64)
        }
    }
    
    private static func runFile(path: String) throws {
        let contents = try String(contentsOfFile: path, encoding: .utf8)
        try run(source: contents)
        if hadError {
            exit(65)
        }
    }
    
    private static func runPrompt() throws {
        let printPrompt = { print("> ", terminator: "") }
        printPrompt()
        while let line = readLine() {
            try run(source: line)
            hadError = false
            printPrompt()
        }
    }
    
    private static func run(source: String) throws {
        let scanner = Scanner(source: source)
        let tokens = scanner.scanTokens()
        for token in tokens {
            print(token)
        }
    }
    
    // TODO: Consider refactoring
    static func error(line: Int, message: String) {
        report(line: line, at: "", message: message)
    }
    
    static func report(line: Int, at: String, message: String) {
        print("[line \(line)] Error\(at): \(message)", to: &standardError)
        hadError = true
    }
}
