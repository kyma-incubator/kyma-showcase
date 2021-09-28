import { EmotionsTemplate, EmotionBar } from './EmotionsCard.styles';

const EmotionsCard = ({ emotionsArray }) => {
  return (
    <EmotionsTemplate>
      {emotionsArray.map((obj, i) => {
        return (
          <div key={i}>
            <p key={obj}>{obj.emotion}</p>
            {obj.emotionValue === 0 ? <p key={i}>Unknown</p> : <EmotionBar emotionValue={obj.emotionValue} key={i} />}
          </div>
        );
      })}
    </EmotionsTemplate>
  );
};

export default EmotionsCard;
