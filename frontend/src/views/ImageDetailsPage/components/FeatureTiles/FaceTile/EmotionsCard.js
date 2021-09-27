import { EmotionsTemplate, EmotionBar } from './EmotionsCard.styles';

const EmotionsCard = ({ emotionsArray }) => {
  return (
    <EmotionsTemplate>
      {emotionsArray.map((obj) => {
        return (
          <>
            <p>{obj.emotion}</p>
            {obj.emotionValue === 'Unknown' ? <p>'Unknown'</p> : <EmotionBar emotionValue={obj.emotionValue} />}
          </>
        );
      })}
    </EmotionsTemplate>
  );
};

export default EmotionsCard;
