var wakeuptime = 7;
var noon = 12;
var lunchtime = 12;
var naptime = lunchtime + 2;
var partytime;
var evening = 18;

// Getting it to show the current time on the page
var showCurrentTime = function()
{
    // display the string on the webpage
    var clock = document.getElementById('clock');
 
    var currentTime = new Date();
 
    var hours = currentTime.getHours();
    var minutes = currentTime.getMinutes();
    var seconds = currentTime.getSeconds();
    var meridian = "AM";
 
    // Set hours
	  if (hours >= noon)
	  {
		  meridian = "PM";
	  }

	  if (hours > noon)
	  {
		  hours = hours - 12;
	  }
 
    // Set Minutes
    if (minutes < 10)
    {
        minutes = "0" + minutes;
    }
 
    // Set Seconds
    if (seconds < 10)
    {
        seconds = "0" + seconds;
    }
 
    // put together the string that displays the time
    var clockTime = hours + ':' + minutes + ':' + seconds + " " + meridian + "!";
 
    clock.innerText = clockTime;
};

// Getting the clock to increment on its own and change out messages and pictures
var updateClock = function() 
{
  var time = new Date().getHours();
  var messageText;
  var image = "https://joincake.imgix.net/steve-johnson-TnjzgonRFPU-unsplash.jpg";

  var timeEventJS = document.getElementById("timeEvent");
  var lolcatImageJS = document.getElementById('lolcatImage');
  
  if (time == partytime)
  {
    image = "https://media.istockphoto.com/photos/nicelooking-attractive-gorgeous-glamorous-elegant-stylish-cheerful-picture-id1165055006?k=20&m=1165055006&s=612x612&w=0&h=OD4-_BceL_R2eaaBzDQrXNIyydwYXOJX-m-0z12z17s=";
    messageText = "Let's party!";
  }
  else if (time == wakeuptime)
  {
    image = "https://img.freepik.com/free-vector/wake-up-fresh-new-day-bedroom-concept_225067-72.jpg?size=338&ext=jpg";
    messageText = "Wake up!";
  }
  else if (time == lunchtime)
  {
    image = "https://pbs.twimg.com/media/Eon-YoqWEAE2x_E.jpg";
    messageText = "Let's have some lunch!";
  }
  else if (time == naptime)
  {
    image = "https://www.canr.msu.edu/contentAsset/image/3f3355d0-6e07-4268-9490-09d2d6cd2cff/fileAsset/filter/Resize,Jpeg/resize_w/750/jpeg_q/80";
    messageText = "Sleep tight!";
  }
  else if (time < noon)
  {
    image = "https://www.poynter.org/wp-content/uploads/2019/07/shutterstock_264132746.jpg";
    messageText = "Good morning!";
  }
  else if (time >= evening)
  {
    image = "http://wallup.net/wp-content/uploads/2017/11/17/285055-night-Sun.jpg";
    messageText = "Good evening!";
  }
  else
  {
    image = "https://media.istockphoto.com/photos/the-ocean-breeze-brings-a-life-of-ease-picture-id1133819651?b=1&k=20&m=1133819651&s=170667a&w=0&h=kjIA9jOXwqRTq0GPWJUR3iISDzUuWGsYwQT3RZ8RkDQ=";
    messageText = "Good afternoon!";
  }

  console.log(messageText); 
  timeEventJS.innerText = messageText;
  lolcatImage.src = image;
  
  showCurrentTime();
};
updateClock();

// Getting the clock to increment once a second
var oneSecond = 1000;
setInterval( updateClock, oneSecond);


// Getting the Party Time Button To Work
var partyButton = document.getElementById("partyTimeButton");

var partyEvent = function()
{
    if (partytime < 0) 
    {
        partytime = new Date().getHours();
        partyTimeButton.innerText = "Party Over!";
        partyTimeButton.style.backgroundColor = "#0A8DAB";
    }
    else
    {
        partytime = -1;
        partyTimeButton.innerText = "Party Time!";
        partyTimeButton.style.backgroundColor = "#222";
    }
};

partyButton.addEventListener("click", partyEvent);
partyEvent(); 


// Activates Wake-Up selector
var wakeUpTimeSelector =  document.getElementById("wakeUpTimeSelector");

var wakeUpEvent = function()
{
    wakeuptime = wakeUpTimeSelector.value;
};

wakeUpTimeSelector.addEventListener("change", wakeUpEvent);


// Activates Lunch selector
var lunchTimeSelector =  document.getElementById("lunchTimeSelector");

var lunchEvent = function()
{
    lunchtime = lunchTimeSelector.value;
};

lunchTimeSelector.addEventListener("change", lunchEvent);


// Activates Nap-Time selector
var napTimeSelector =  document.getElementById("napTimeSelector");

var napEvent = function()
{
    naptime = napTimeSelector.value;
};

napTimeSelector.addEventListener("change", napEvent);