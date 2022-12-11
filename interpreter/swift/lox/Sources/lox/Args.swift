import Foundation

enum Args {
    case unrecognized([String])
    case help
    case prompt
    case run(String)
    case format(String)
}

extension Args {
    static func get() -> Args {
        let args = Array(CommandLine.arguments.dropFirst())
        
        switch args.count {
        case 0:
            return .prompt
        case 1:
            if args[0] == "help" {
                return .help
            }
        case 2:
            switch args[0] {
            case "format":
                return .format(args[1])
            case "run":
                return .run(args[1])
            default:
                break
            }
        default:
            break
        }
        
        return unrecognized(args)
    }
    
    static func usage() -> String {
        let programName = "lox"
        var s = "Usage:\n"
        s += "  \(programName)                    runs REPL\n"
        s += "  \(programName) run [script]       runs the script\n"
        s += "  \(programName) format [script]    formats the script\n"
        s += "  \(programName) help\n"
        return s
    }
}
