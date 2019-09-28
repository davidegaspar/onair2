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
            track: resp.data[0].track.title,
            author: resp.data[0].track.artist,
          });
        } else {
          render({
            playlist: null
          });
        }
      }else{
        render({
          playlist: null
        });
      }
    }
    request.onerror = function() {
      render({
        playlist: null
      });
    }
    request.send();
  }

  function render(state){
    console.log(state);

    if (!state.playlist) {
      //onair
      try {
        document.getElementById('centova-onair').innerText = 'OFFLINE'
        document.getElementById('centova-title').innerText = ''
        document.getElementById('centova-title').href = ''
        document.getElementById('centova-host').innerText = ''
        document.getElementById('centova-from').innerText = ''
        document.getElementById('centova-to').innerText = ''
        document.getElementById('centova-img').style.backgroundImage = ''
      } catch (e) {
        console.log('WARN: onair not found')
      }
    } else {
      //onair
      try {
        document.getElementById('centova-onair').innerText = 'ON AIR'
        document.getElementById('centova-title').innerText = centova.data[state.playlist].title
        document.getElementById('centova-title').href = centova.data[state.playlist].link
        document.getElementById('centova-host').innerText = centova.data[state.playlist].host
        document.getElementById('centova-from').innerText = centova.data[state.playlist].from
        document.getElementById('centova-to').innerText = centova.data[state.playlist].to
        document.getElementById('centova-img').style.backgroundImage = 'url(' + centova.data[state.playlist].img + ')'
      } catch (e) {
        console.log('WARN: onair not found')
      }

      // player
      try {
        document.getElementById('qtradiotitle').innerText = centova.data[state.playlist].title
        document.getElementById('qtFeedPlayerTrack').innerText = state.track
        document.getElementById('qtFeedPlayerAuthor').innerText = state.author
      } catch (e) {
        console.log('WARN: onair not found')
      }
    }
  }

  get();
  setInterval(get, centova.refresh_in_seconds * 1000);

})(window.centova)
