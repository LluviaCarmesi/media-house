namespace media_house_api.Models
{
    public class Show
    {
        private int id = 0;
        private string title = string.Empty;

        public Show(int id, string title)
        {
            this.id = id;
            this.title = title;
        }

        public int Id
        {
            get
            {
                return id;
            }
            set
            {
                id = value;
            }
        }

        public string Title
        {
            get
            {
                return title;
            }
            set
            {
                title = value;
            }
        }
    }
}
