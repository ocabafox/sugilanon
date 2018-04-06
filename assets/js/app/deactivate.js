$('#deactivate-modal').on({
  'hide.uk.modal': function() {
    $('#deactivate').trigger('reset');
    $('input[type="submit"]').prop('disabled', true);
  }
});

function validate(appUsername) {
  if (appUsername === $('input[name="username"]').val()) {
    $('input[type="submit"]').prop('disabled', false);
  } else {
    $('input[type="submit"]').prop('disabled', true);
  }
}
