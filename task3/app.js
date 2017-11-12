$(document).ready(function() {
	var tag = $("#root");
	tag.append($("<ul>"));
	tag = tag.children[0];
	tag.append($("<li>"));
	tag = tag.children[0];
	tag.text("<span> “Сделать задание #3 по web-программированию” </span>");

})