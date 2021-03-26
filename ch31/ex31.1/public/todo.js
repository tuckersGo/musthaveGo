//ch31/ex31.1/public/todo.js
(function($) {
    'use strict';
    $(function() {
        var todoListItem = $('.todo-list');             // ❶ 항목이 위치하는 곳
        var todoListInput = $('.todo-list-input');      // ❷ 할 일 제목 입력 박스
    
        $('.todo-list-add-btn').on("click", function(event) {
            event.preventDefault();                     // ❸ add 버튼 클릭시
    
            var item = $(this).prevAll('.todo-list-input').val();
    
            if (item) {
                $.post("/todos", JSON.stringify({name:item}), addItem); // ❹ POST 요청
                todoListInput.val("");
            }
        });
    
        var addItem = function(item) {        // ➎ 항목 추가 함수
            if (item.completed) {
                todoListItem.append("<li class='completed'"+ " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            } else {
                todoListItem.append("<li "+ " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            }
        };
    
        $.get('/todos', function(items) {     // ➏ 웹 페이지가 뜰때 GET 요청 전송
            items.forEach(e => {
                addItem(e)
            });
        });
    
        todoListItem.on('change', '.checkbox', function() {
            // ➐ 완료 체크 박스 클릭시
            var id = parseInt($(this).closest("li").attr('id'));
            var name = $(this).closest("li").text();
            var $self = $(this);
            var complete = true;
            if ($(this).attr('checked')) {
                complete = false;
            }
            // ➑ todos/id로 PUT 요청 전송
            $.ajax({
                url: "/todos/" + id,
                type: "PUT",
                data: JSON.stringify({id:id, name:name, completed:complete}),
                success: function(data) {
                    if (complete) {
                        $self.attr('checked', 'checked');
                    } else {
                        $self.removeAttr('checked');
                    }
            
                    $self.closest("li").toggleClass('completed');
                }
            })
        });
    
    
        // ➒ 삭제 버튼 클릭 시DELETE 요청 전송
        todoListItem.on('click', '.remove', function() {
            // url: todos/id method: DELETE
            var id = $(this).closest("li").attr('id');
            var $self = $(this);
            $.ajax({
                url: "/todos/" + id,
                type: "DELETE",
                success: function(data) {
                    if (data.success) {
                        $self.parent().remove();
                    }
                }
            })
        });
    
    });
})(jQuery);