
async function Get_Data(url, dataArray) {
    try {
        const response = await fetch(url);
        const dt = await response.json();
        dataArray.push(dt);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

async function Get_All_Posts() {
    let postsData = [];
    let res = "";
    await Get_Data("https://jsonplaceholder.typicode.com/posts", postsData);
    postsData = postsData[0]
        for (let i = 0; i < postsData.length; i++) {
            res += `
            <div>
                <div>${postsData[i].title}</div><br><br>
                <div>${postsData[i].body}</div>
                

            </div>
            `;

            
            
        }
        document.getElementById("posts").innerHTML = res
        
    }

    
    Get_All_Posts();

    const trying = async () => {
        try {
            let f = await fetch("/testing")
            let r = await f.json()
            console.log(r)
            
        }catch(err) {
            console.error(err)
        }

    }
    trying()