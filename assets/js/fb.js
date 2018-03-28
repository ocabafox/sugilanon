window.fbAsyncInit = function() {
  FB.init({
    appId      : '1618996761552073',
    cookie     : true,
    xfbml      : true,
    version    : 'v2.12'
  });

  FB.getLoginStatus(function(response) {
    statusChangeCallback(response);
  });
};

function statusChangeCallback(response) {
  if (response.status === 'connected') {
    FB.api('/me?fields=id,name,email', function(response) {
      if (response && !response.error) {
	console.log(response);
      }
    })
  }
}

function checkLoginState() {
  FB.getLoginStatus(function(response) {
    if (response.status === 'connected') {
      console.log($);
      statusChangeCallback(response);
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
