function submitStory(evt) {
  evt.preventDefault();

  $.post('/story', {
    owner: $('.story-owner').val(),
    title: $('.story-title').val(),
    body: $('.story-body').val(),
  }).done(function(res) {
    if (res.status === 'success') {
      console.log('DISPLAY SUCCESS MESSAGE');
    } else {
      console.log('DISPLAY ERROR MESSAGE');
    }
  }).always(function() {
    $('.story-form').trigger('reset');
  });
}
