var allDays = document.querySelectorAll('.qt-show-schedule .tabs .tab a');
var today = new Date().getDay();
var dayOfTheWeek = today === 0 ? 6 : today-1;
allDays[dayOfTheWeek].click();
