
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
    await Get_Data("https://dummyjson.com/posts", postsData);
    postsData = postsData[0].posts
        for (let i = 0; i < postsData.length; i++) {
            
            res += `
            <div class="card" onclick="openAlert(${i})">
                <div class="card-title">${postsData[i].title || "No Title"}</div>
                <div class="card-body">${postsData[i].body}</div><br> 
                <div class="card-tags">${postsData[i].tags.join(', ') || "No Tags"}</div>
                 <div class="card-like-dislike">
                    <b>⬆</b>${postsData[i].reactions.likes || 0} | 
                    <b>⬇</b>${postsData[i].reactions.dislikes || 0}
                    </div> 
            </div>
            `;

            
            
        }
        document.getElementById("posts").innerHTML = res
        
        
    }

    
    Get_All_Posts();

    // const trying = async () => {
    //     try {
    //         let f = await fetch("/testing")
    //         let r = await f.json()
    //         console.log(r)
            
    //     }catch(err) {
    //         console.error(err)
    //     }

    // }
    // trying()