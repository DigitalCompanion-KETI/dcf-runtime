import mesh

class Handler:
    def __init__(self):
        pass

    def __call__(self, req):

        # mesh call
        #
        # functionName = "<FUNCTIONNAME>"
        # input = req.input
        # result = mesh.mesh_call(functionName, input)
        # return result

        # single call
        return req.input
