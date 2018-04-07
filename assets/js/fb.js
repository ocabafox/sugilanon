function login() {
  FB.login(function(response) {
    if (response.status === 'connected') {
      FB.api('/me?fields=id,name,email,link,gender,website,updated_time', function(response) {
        if (response && !response.error) {
          $.post("/login", {
            facebook_id: response.id,
            name: response.name,
            email: response.email,
            link: response.link,
            gender: response.gender,
            website: response.website,
            updated: response.updated_time
          }).done(function() {
            location.reload();
          });
        }
      })
    }
  });
}

function logout() {
  FB.getLoginStatus(function(response) {
    if (response.status === 'connected') {
      FB.logout(function(response) {
        $.get("/logout").done(function() {
          location.reload();
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
