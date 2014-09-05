import tornado.ioloop
import tornado.web

class UploadFile(tornado.web.RequestHandler):
    def post(self):
        print self.request.files['data'][0]['filename']

application = tornado.web.Application([
    (r"/", UploadFile)
])

if __name__ == "__main__":
    application.listen(6666)
    tornado.ioloop.IOLoop.instance().start()
