import CalendarComp from "../components/calendario/CalendarComp";
import images from "../components/rooms/room-Image";
import ImageGallery from "react-image-gallery";
import Navbar from "../components/navbar/Navbar";
import "./hoteles.css";

function App() {
  return (
    <>
      <Navbar />
      <div className="ContainerImagenes">
        <ImageGallery
          items={images}
          showThumbnails={false}
          showBullets={true}
          slideInterval={3000}
        />
      </div>
      <div className="info">
        <CalendarComp />
      </div>
    </>
  );
}

export default App;
