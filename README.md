1. # ByteDance

   ## Process

   - 基础功能实现

   | 功能       | 是否完成 | 是否存在BUG |
   | ---------- | -------- | :---------: |
   | 视频Feed流 | 完成     |  暂未发现   |
   | 视频投稿   | 完成     |  暂未发现   |
   | 个人主页   | 完成     |  暂未发现   |
   | 喜欢列表   | 完成     |  暂未发现   |
   | 用户评论   | 完成     |  暂未发现   |

   ## Structure

   ```
   .
   ├── LICENSE
   ├── README.md
   ├── api          // 运行需要的接口
   ├── conf    	 // 配置文件以及mysql表结构
   ├── controller   // Controller 层业务逻辑
   ├── dal          // DAl 层业务逻辑
   ├── main.go      // 主函数
   ├── middleware   // 中间件
   ├── router       // 路由设置
   ├── service      // Service 层业务逻辑
   └── utils        // 一些工具函数
   ```
   ## Getting started
   
   1. Install `ffmpeg` and add it to your `PATH`.
   
       We use `ffmpeg` just to extract the video frame as the cover image.
   
       Please refer to some documents or search in google to find `How to install ffmpeg?`.
       
       If you can't install it, it doesn't matter.  And we just simply use the `default.jpg` instead.

   2. Change `server.host` in `conf.yaml` to your **Local Machine's IP Address**.

       Please refer to some documents or search in google to find `How to get your computer's IP Address?`
   
   3. Set `BaseUrl` in the app.

       Refer to [https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#mC5eiD](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7#mC5eiD).
       
       If you successfully get your IP Address, please use **your own IP Address** instead `http://192.168.1.7:8080` in the above document to fill the `BaseUrl` field in the advanced setting.

       For example, if your own IP Addresss is `192.168.108`, you should fill the `BaseUrl` to `http://192.168.108:8080` (`8080` is the port we set in the `conf.yaml`).

   4. Just run this command in the root folder of your project:
       ```
       $ go run main.go
       ```
   
   5. Can also run this command in the root folder of your project:
   
   ```
   $ go build
   $ ./DouyinSimpleProject
   ```
   
   ## Note
   
   ### 视频流
   
   手机端第一次打开首页时，会按照当前的时间向服务端发起请求，获取 `limitNum` 个视频。

   当用户向下滑动到第 `limitNum-2` 个（即倒数第三个）, 手机端就会把第 `limitNum` 个视频的创建时间作为 `next_time` 向服务端发送新的请求, 获取新的 `limitNum` 个视频列表，以此来保证用户体验。

   因此在测试的时候我们将 `limitNum` 设为 5，即当你刷到第三个视频的时候，手机端就会立即向后端发送请求，以此来获取下一批次的五个视频。

   当刷到数据库中的最后一个视频的时候，会从头开始播放。

   ------
   
   
