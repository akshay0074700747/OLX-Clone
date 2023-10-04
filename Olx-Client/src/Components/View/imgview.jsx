import "./View.css";

export default function ImageView({ item }) {
  var ii;
  return (
    <div>
      {item.images.map((imageview) => {
        const base64ImageData = `data:image/png;base64,${imageview.image}`;
        return (
          <>
            <div className="imageShowDiv">
              <img src={base64ImageData} alt="" />
            </div>
            {ii++}
            {ii % 2 == 0 && <br></br>}
          </>
        );
      })}
    </div>
  );
}
