## BEM (Block Element Modifier)

A structured way of naming classes that is based on the properties of the element


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