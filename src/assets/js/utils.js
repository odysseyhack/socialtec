/*jshint esversion: 8 */

promisify= (inner) =>
    new Promise((resolve, reject) =>
        inner((err, res) => {
            if (err) {
                reject(err);
            } else {
                resolve(res);
            }
        })
    );

function in_my_interest(offers) {
  my_interests = Storage.get_my_interests();
  common = my_interests.filter(value => -1 !== offers.indexOf(value));
  return common;
}

// TODO: Make Popup GUI responsive
function Popup(offer_name, offer_details, interest_name, interest_details) {
  $("#offer_name").text(offer_name);
  $("#offer_details").text(offer_details);
  $("#interest_name").text(interest_name);
  $("#interest_details").text(interest_details);
  $("#match-found").show();
}
