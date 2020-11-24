from flask import Flask, request, jsonify


app=Flask(__name__)

@app.route('/api/search/',methods=['GET','POST'])
def search_author_by_name():
    from scholarly import scholarly
    author_name='Mark'
    res=next(scholarly.search_author(author_name)).affiliation
    return jsonify(res)

if __name__ == '__main__':
    app.run(host="127.0.0.1", port=8080, debug=True)

