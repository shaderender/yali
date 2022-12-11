class Scanner {
    fileprivate var source: String
    fileprivate var tokens: [Token] = []
    
    fileprivate var start: Int = 0
    fileprivate var current: Int = 0
    fileprivate var line: Int = 1
    
    init(source: String) {
        self.source = source
    }
    
    func scanTokens() -> [Token] {
        while !isAtEnd() {
            start = current
            scanToken()
        }
        tokens.append(Token(type: .eof, line: line))
        return tokens
    }
    
    func scanToken() {
        guard let c = advance() else { return }
        
        switch c {
        // Single-character tokens.
        case "(":
            addToken(type: .leftParen)
        case ")":
            addToken(type: .rightParen)
        case "{":
            addToken(type: .leftBrace)
        case "}":
            addToken(type: .rightBrace)
        case ",":
            addToken(type: .comma)
        case ".":
            addToken(type: .dot)
        case "-":
            addToken(type: .minus)
        case "+":
            addToken(type: .plus)
        case ";":
            addToken(type: .semicolon)
        case "*":
            addToken(type: .star)
        // Operators
        case "!":
            addToken(type: match("=") ? .bangEqual : .bang)
        case "=":
            addToken(type: match("=") ? .equalEqual : .equal)
        case "<":
            addToken(type: match("=") ? .lessEqual : .less)
        case ">":
            addToken(type: match("=") ? .greaterEqual : .greater)
        // Longer lexemes
        case "/":
            if match ("/") {
                comment()
            } else {
                addToken(type: .slash)
            }
        case " ", "\r", "\t":
            break // Ignore whitespace
        case "\n":
            addToken(type: .newline)
            line += 1
        case "\"":
            string()
        default:
            if c.isNumber {
                number()
            } else if c.isLetter {
                identifier()
            } else {
                Lox.error(line: line, message: "Unexpected character '\(c)'.")
            }
        }
    }
    
    // TODO: Can we remove this and rely on optionals from functions like advance() ?
    private func isAtEnd() -> Bool {
        return current >= source.count
    }
    
    private func characterAt(_ offset: Int) -> Character? {
        if offset < 0 || offset >= source.count {
            return nil
        }
        let i = source.index(source.startIndex, offsetBy: offset)
        return source[i]
    }
    
    // Half-open range
    private func stringAt(start: Int, to end: Int) -> String? {
        guard start >= 0 && start < source.count,
              end > 0 && end <= source.count,
              start <= end else {
            return nil
        }
        
        let i = source.index(source.startIndex, offsetBy: start)
        let j = source.index(source.startIndex, offsetBy: end)
        return String(source[i..<j])
    }
    
    private func peek() -> Character? {
        return characterAt(current)
    }
    
    private func peekTwo() -> Character? {
        return characterAt(current + 1)
    }
    
    // Modifying
    private func advance() -> Character? {
        guard current < source.count else { return nil }
        let i = source.index(source.startIndex, offsetBy: current)
        current += 1
        return source[i]
    }
    
    private func match(_ c: Character) -> Bool {
        guard let cur = characterAt(current),
              c == cur else { return false }
        current += 1
        return true
    }
    
    private func addToken(type: TokenType) {
        tokens.append(Token(type: type, line: line))
    }
    
    private func string() {
        while let next = peek(), next != "\"" {
            if next == "\n" {
                line += 1
            }
            _ = advance()
        }
        
        if isAtEnd() {
            Lox.error(line: line, message: "Unterminated string.")
            return
        }
        
        // The closing ".
        _ = advance()
        
        // Trim the surrounding quotes
        guard let text = stringAt(start: start + 1, to: current - 1) else {
            fatalError()
        }
        addToken(type: .string(text))
    }
    
    private func number() {
        let advanceNumbers = {
            // TODO: This may produce some odd results, since Character.isNumber works on fractions and the like.
            while let next = self.peek(), next.isNumber {
                _ = self.advance()
            }
        }
        
        advanceNumbers()
        
        if let next = peek(), next == ".",
           let nextNext = peekTwo(), nextNext.isNumber {
            _ = advance()
            
            advanceNumbers()
        }
        
        guard let text = stringAt(start: start, to: current),
              let value = Double(text) else {
            fatalError()
        }
        
        addToken(type: .number(value))
    }
    
    private func identifier() {
        let isAlphaNumeric = { c in
            return c.isUppercase || c.isLowercase || c == "_" || c.isNumber
        }
        
        while let next = peek(), isAlphaNumeric(next) {
            _ = advance()
        }
        
        guard let text = stringAt(start: start, to: current) else {
            fatalError()
        }
        
        if let keyword = TokenType.keywords[text] {
            addToken(type: keyword)
        } else {
            addToken(type: .identifier(text))
        }
    }
    
    private func comment() {
        // A commment goes until the end of the line.
        while let next = peek(), next != "\n" {
            _ = advance()
        }
        
        guard let text = stringAt(start: start, to: current) else {
            fatalError()
        }
        
        addToken(type: .comment(text))
    }
}
