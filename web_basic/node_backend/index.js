import express from 'express'

const port = 8000

const app = express()
app.use(express.json())
app.use(express.urlencoded({extended: true}))

app.get('/', (req, res) =>{
    res.status(200).send("Welcome to js backend")
})

app.get('/get', (req, res) =>{
    res.status(200).send("Welcome to get route")
})
app.post('/post',(req, res)=>{
    res.status(200).send(req.body)
})

app.post('/postform', (req, res)=>{
    res.send(JSON.stringify(req.body)).status(200)
})

app.listen(port, ()=>{
    console.log(`Server is running at ${port}`)
})