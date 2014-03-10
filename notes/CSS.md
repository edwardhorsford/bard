# CSS

## Normal Flow
Boxes in the normal flow belong to a formatting context, which may be block or inline, but not both simultaneously. Block boxes participate in a block formatting context. Inline boxes participate in an inline formatting context.

## SMACSS (Scalable/Modular Architecture for CSS)

## BEM (Block Element Modifier)
```
.signup_form-email {
	/* styling for email input */
}
```
A structured way of naming classes that is based on the properties of the element


## Positioning

### Static

### Relative

### Absolute

### Fixed

### inherit

## Box-sizing
```
*, *:before, *:after {
	-moz-box-sizing: border-box;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
}
```

## Clearfix

```
.clearfix:before, .clearfix:after {
	content: ' ';
	display: table;
}

.clearfix { zoom: 1; }
```