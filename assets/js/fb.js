window.fbAsyncInit = function() {
  FB.init({
    appId: $('#app-id').data('app-id'),
    cookie: true,
    xfbml: true,
    version: 'v2.12'
  });
};

function login() {
  FB.login(function(response) {
    if (response.status === 'connected') {
      FB.api('/me?fields=id,name,email,link,gender,updated_time', function(response) {
        if (response && !response.error) {
          $.post('/login', {
            facebook_id: response.id,
            name: response.name,
            email: response.email,
            link: response.link,
            gender: response.gender,
            updated: response.updated_time
          }).done(function() {
            window.location.replace('/');
          });
        }
      })
    }
  }, {
    scope: 'email, public_profile'
  });
}

function logout() {
  FB.getLoginStatus(function(response) {
    if (response.status === 'connected') {
      FB.logout(function(response) {
        $.get('/logout').done(function() {
          window.location.replace('/');
        });
      });
    }
  });
}

(function(d, s, id){
  var js, fjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) {return;}
  js = d.createElement(s); js.id = id;
  js.src = "https://connect.facebook.net/en_US/sdk.js";
  fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));
