# http-to-https
This program reverse-proxies any process running over http on your local machine to an https address on the public internet. 

I created it as a way to replace [ngrok](https://ngrok.com) for my personal use. Unlike ngrok, you can make as many https proxies as you want just with repeated calls to `MakeProxy()`. For a full description see [my blog post](https://jchavez.io/posts/https).

<p align="center">
  <kbd>
    <a href="https://jchavez.io/posts/https">
      <img src="https://i.postimg.cc/pLfbMPBC/Screenshot-2020-09-06-at-12-26-00-PM.png" alt="blog" width="600vw"/>
    </a>
  </kbd>
</p>
