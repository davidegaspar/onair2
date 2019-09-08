;(function(centova) {

  function get() {
    let request = new XMLHttpRequest();
    request.open('GET', centova.url, true);
    request.onload = function() {
      if (request.status >= 200 && request.status < 400) {
        let resp=JSON.parse(request.responseText);//console.log(resp);
        if (resp.hasOwnProperty('data') && resp.data.length > 0 && resp.data[0].offline === false) {
          render({
            playlist: resp.data[0].track.hasOwnProperty('playlist') ? resp.data[0].track.playlist.id + '' : null,
          });
        } else {
          render({
            playlist: 'OFF'
          });
        }
      }else{
        render({
          playlist: 'OFF'
        });
      }
    }
    request.onerror = function() {
      render({
        playlist: 'OFF'
      });
    }
    request.send();
  }

  function render(state){
    console.log(state);
    if (state.playlist === 'OFF') {
      document.getElementById('centova-onair').innerText = 'OFF AIR'
    } else {
      document.getElementById('centova-onair').innerText = 'ON AIR'
    }
    document.getElementById('centova-title').innerText = centova.data[state.playlist].title
    document.getElementById('centova-title').href = centova.data[state.playlist].link
    document.getElementById('centova-host').innerText = centova.data[state.playlist].host
    document.getElementById('centova-from').innerText = centova.data[state.playlist].from
    document.getElementById('centova-to').innerText = centova.data[state.playlist].to
    document.getElementById('centova-img').style.backgroundImage = 'url(' + centova.data[state.playlist].img + ')'
  }

  get();
  setInterval(get, centova.refresh_in_seconds * 1000);

})(window.centova)
