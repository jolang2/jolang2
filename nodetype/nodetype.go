package nodetype

type NodeType string

const (
	PACKAGE_DECLARATION     = NodeType("package_declaration")
	CLASS_DECLARATION       = NodeType("class_declaration")
	CONSTRUCTOR_DECLARATION = NodeType("constructor_declaration")
	METHOD_DECLARATION      = NodeType("method_declaration")
	CLASS_BODY              = NodeType("class_body")
	IMPORT_DECLARATION      = NodeType("import_declaration")
	IDENTIFIER              = NodeType("identifier")
	MODIFIERS               = NodeType("modifiers")
	FIELD_DECLARATION       = NodeType("field_declaration")
	VARIABLE_DECLARATOR     = NodeType("variable_declarator")
	EQUAL                   = NodeType("=")
	STATIC                  = NodeType("static")
	FORMAL_PARAMETERS       = NodeType("formal_parameters")
	FORMAL_PARAMETER        = NodeType("formal_parameter")
)

func (n NodeType) String() string {
	return string(n)
}
