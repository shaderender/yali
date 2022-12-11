enum TokenType {
    // Single-character tokens.
    case leftParen, rightParen, leftBrace, rightBrace, comma, dot, minus, plus, semicolon, slash, star
    
    // One or two character tokens.
    case bang, bangEqual, equal, equalEqual, greater, greaterEqual, less, lessEqual
    
    // Literals.
    case identifier(String)
    case string(String)
    case number(Double)
    case comment(String)
    
    // Keywords.
    case kAnd, kClass, kElse, kFalse, kFun, kFor, kIf, kNil, kOr, kPrint, kReturn, kSuper, kThis, kTrue, kVar, kWhile
    
    case newline
    case eof
    
    static let keywords: [String:TokenType] = [
        "and":    .kAnd,
        "class":  .kClass,
        "else":   .kElse,
        "false":  .kFalse,
        "fun":    .kFun,
        "for":    .kFor,
        "if":     .kIf,
        "nil":    .kNil,
        "or":     .kOr,
        "print":  .kPrint,
        "return": .kReturn,
        "super":  .kSuper,
        "this":   .kThis,
        "true":   .kTrue,
        "var":    .kVar,
        "while":  .kWhile,
    ]
}
