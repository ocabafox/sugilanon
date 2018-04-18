function submitStory(evt) {
  evt.preventDefault();
  
	var title = $('.story-title').val();
	var body = $('.story-body').val();
	
	if (title === "" && body === "") {
		return;
	}

  $.post('/story', {
    owner: $('.story-owner').val(),
    title: title,
    body: body
  }).done(function(res) {
    if (res.status === 'success') {
    	$('.uk-label').html(res.message);
    } else {
			$('.uk-label').removeClass('uk-label-success').addClass('uk-label-warning').html(res.message);
    }
  }).always(function() {
    $('.story-form').trigger('reset');
  });
}
