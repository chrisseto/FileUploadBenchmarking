import tornado.ioloop
import tornado.web

class UploadFile(tornado.web.RequestHandler):
    def post(self):
        print(self.request.files['file'][0]['filename'])

application = tornado.web.Application([
    (r"/", UploadFile)
])

