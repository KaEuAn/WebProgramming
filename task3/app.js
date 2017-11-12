$(document).ready(function() {
	var tag = $("#root");
	tag.append("<ul>");
	tag = tag.children([0]);
	tag.attr("id", "ul");
	add_li("Сделать задание #3 по web-программированию");
	$(root).append("<input id='add_task_input'>");
	$(root).append($("<button id='add_task' value='add'>").text('add_task'))


	$("#add_task").click(function() {
		add_li($("#add_task_input").val());
	})
})

function add_li(string) {
	var ultag = $("ul");
	ultag.append("<li>");
	ultag = ultag.children().last();
	ultag.append("<span>");
	ultag.children([0]).text(string);
	ultag.append($("<button class='delete'>").text("Удалить"))
	$(".delete").click(function () {
		$(this).parent().remove();
	})

}