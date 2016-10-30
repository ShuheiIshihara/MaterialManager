$(function(){
    $('#kancolle-btnRegist').on('click', function(){

        // 資源
        var data = $('#kancolle-jsonData').val();
        data = data.replace(/svdata=/g, "");
        // ランキング
        var ranking = $('#kancolle-ranking').val();
        // 戦果
        var senka = $('#kancolle-senka').val();
        // 日付
        var date = $('#kancolle-date').val();

        // 入力チェック
        if(!isData(data, ranking, senka)){
            return;
        }

        var jsonData = JSON.parse(data);
        // console.log(jsonData);

        // 資源
        var material = [];
        $.each(jsonData.api_data.api_material, function(index, val){
            console.log(index, val);
            material[index] = val.api_value;
        });
        console.log(material);
        
        // Level
        var level = jsonData.api_data.api_basic.api_level;
        console.log("レベル: ", level);

        // 戦果
        // 出撃の勝利数
        var stWin = jsonData.api_data.api_basic.api_st_win;
        // 出撃の敗北数
        var stLose = jsonData.api_data.api_basic.api_st_lose;

        console.log(stWin, stLose);

        // 遠征の回数
        var msCnt = jsonData.api_data.api_basic.api_ms_count;
        // 遠征の成功数
        var msSucc = jsonData.api_data.api_basic.api_ms_success;

        console.log(msCnt, msSucc);

        // 演習の勝利数
        var ptWin = jsonData.api_data.api_basic.api_pt_win;
        // 演習の敗北数
        var ptLose = jsonData.api_data.api_basic.api_pt_lose;

        console.log(ptWin, ptLose);

        // DB登録依頼
        insert(level, material, stWin, stLose, msCnt, msSucc, ptWin, ptLose, senka, ranking, date);
    });
});

// 入力チェック
function isData(data, ranking, senka){
    if(data === ''){
            alert("資源情報がありません");
            return false;
    }

    if(ranking === '' || 
        isNaN(parseInt(ranking, 10))) {
        alert("ランキング情報に誤りがあります");
        return false;
    }

    if(senka === '' ||
        isNaN(parseInt(senka, 10))) {
        alert("戦果情報に誤りがあります");
        return false;
    }
    return true;
}

// DB登録依頼
function insert(level, material, stWin, stLose, msCnt, msSucc, ptWin, ptLose, senka, ranking, date){
    $.ajax({
        type: 'POST',
        url: '/material/store',
        timeout: 10000,
        dataType: 'json',
        data: {
            'level':      level,
            'fuel':       material[0],
            'ammunition': material[1],
            'steel':      material[2],
            'bauxite':    material[3],
            'bucket':     material[5],
            'banner':     material[4],
            'dMaterial':  material[6],
            'screw':      material[7],
            'winning_sortie':       stWin,
            'defeatting_sortie':    stLose,
            'expedition':           msCnt,
            'successs_expedition':  msSucc,
            'winning_exercises':    ptWin,
            'defeatting_exercises': ptLose,
            'veterans':             senka,
            'ranking':              ranking,
            'date':          date
        }
    }).done(function(response, textStatus, jqXHR) {
       // 成功時処理
        //レスポンスデータはパースされた上でresponseに渡される
     }).fail(function(jqXHR, textStatus, errorThrown ) {
        // 失敗時処理
     }).always(function(data_or_jqXHR, textStatus, jqXHR_or_errorThrown) {
        // doneまたはfail実行後の共通処理
    });
}